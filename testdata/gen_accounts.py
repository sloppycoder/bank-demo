import argparse
import configparser
import datetime
import re
import sys
import time
from mimesis import Generic, random
from cassandra.cluster import Cluster
from cassandra.auth import PlainTextAuthProvider
from cassandra import DriverException

PRODUCTS = [
    ['1001', 'Super Saver Account'],
    ['1002', 'Salary Plus Account'],
    ['2001', 'Student Saver Account'],
    ['2002', 'Everyday Rebate Account']
]

NUM_OF_PRODUCTS = len(PRODUCTS)

all_account_ids = {}


class Balance:
    def __init__(self, amount, type, last_updated):
        self.amount = amount
        self.type = type
        self.credit = False
        self.last_updated = last_updated


def unique_account_id(g):
    # try generate a unique account number for 10 times
    # use telephone to proxy an account number
    for i in range(10):
        if i > 1:
            num = re.sub(r"\+|\-|\(|\)|\.|\s", '', g.person.telephone())
        else:
            # always use this account number for 1st account
            # it it hardcoded in some places, so it must exist
            num = '10001000'

        if num not in all_account_ids:
            all_account_ids[num] = 1
            return num
    return None


def random_product(rand):
    prod_code, prod_name = PRODUCTS[int(rand.uniform(0, NUM_OF_PRODUCTS - 1, 0))]
    return prod_code, prod_name


def gen_random_accounts(sess, n):
    success, fail = 0, 0
    g = Generic('en')
    r = random.Random()
    ts_base = int(time.mktime(datetime.date(2019, 1, 1).timetuple()))

    insert_stmt = sess.prepare('''
    INSERT INTO casa_account(account_id, nickname, prod_code, prod_name, 
                             currency, status, status_last_updated, balances)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    ''')

    for i in range(n):
        account_id = unique_account_id(g)
        prod_code, prod_name = random_product(r)
        balance = r.uniform(1000.0, 5_000_000.0, 4)
        ts = int(r.uniform(0, 15_724_800, 0))  # 3600 * 24 * 182
        update_dt = datetime.datetime.fromtimestamp(ts_base + ts)
        # print(update_dt)
        try:
            sess.execute(insert_stmt,
                         [account_id, g.person.name(), prod_code, prod_name, 'THB', 0, update_dt,
                          {Balance(balance, 0, update_dt), Balance(balance, 1, update_dt)}])
            success += 1
        except DriverException as e:
            print(e)
            fail += 1

    return success, fail


def populate_testdata(n, force_drop):
    sess = create_session()
    sess.execute('use vino9')
    print_db_version(sess)
    if force_drop:
        drop_and_recreate_table(sess)

    success, fail = gen_random_accounts(sess, n)
    with open('ids.txt', 'w') as f:
        f.writelines("%s\n" % id for id in all_account_ids)
    print(f'success={success}, fail={fail}, uniq_id={len(all_account_ids)}, accounts_in_db={accounts_in_db(sess)})')

    sess.cluster.shutdown()


def drop_and_recreate_table(sess):
    # drop existing table and UDT
    sess.execute('DROP TABLE IF EXISTS casa_account')
    sess.execute('DROP TYPE IF EXISTS balance')
    # recreate
    sess.execute(''' 
    CREATE TYPE IF NOT EXISTS balance (
        amount float,
        credit boolean,
        type smallint,
        last_updated timestamp
    )
    ''')
    sess.execute('''
    CREATE TABLE casa_account(
        account_id text,
        nickname text, 
        prod_code text,
        prod_name text,
        currency text,
        status smallint,
        status_last_updated timestamp,
        balances set<frozen<balance>>,
        PRIMARY KEY(account_id)
    )
    ''')


def read_env(env: str) -> dict:
    with open(env, 'r') as f:
        config_string = '[root]\n' + f.read()
    config = configparser.ConfigParser()
    config.read_string(config_string)
    return dict([(k, config.get('root', k)) for k in config.options('root')])


def create_session():
    config = read_env('../cassandra.env')
    if config['instance'] == 'astra':
        cluster = Cluster(
            cloud={'secure_connect_bundle': '../secure-connect-vino9.zip'},
            auth_provider=PlainTextAuthProvider(config['username'], config['password']))
        return cluster.connect()
    elif config['instance'] == 'local':
        cluster = Cluster([config['host']], port=int(config['port']))
        return cluster.connect('vino9')
    else:
        print('unknown cassandra instance ${config["instance"]}, abort.')
        sys.exit(1)


def print_db_version(sess):
    row = sess.execute("select release_version from system.local").one()
    if row:
        print('connected to cassandra version ', row[0])
    else:
        print('something is wrong with the session')


def accounts_in_db(sess):
    row = sess.execute("select count(*) from casa_account").one()
    if row:
        return row[0]
    else:
        return None


def init_argparse():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-v", "--version", action="version",
        version=f"{parser.prog} version 1.0.0"
    )
    parser.add_argument('files', nargs='*')
    return parser


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--drop', action='store_true')
    parser.add_argument('n', nargs='?', default=1, type=int)
    args = parser.parse_args()

    print(args)
    populate_testdata(args.n, args.drop)
