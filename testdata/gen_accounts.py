import argparse
import datetime
import re
import sys
import time
import mysql.connector
from mimesis import Generic, random

PRODUCTS = [
    ['1001', 'Super Saver Account'],
    ['1002', 'Salary Plus Account'],
    ['2001', 'Student Saver Account'],
    ['2002', 'Everyday Rebate Account']
]

NUM_OF_PRODUCTS = len(PRODUCTS)

all_account_ids = {}


db_conf = {
    'user': 'demo',
    'password': 'demo',
    'host': '192.168.39.1',
    'database': 'demo',
    'raise_on_warnings': True
}

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
            num = re.sub(r'\+|\-|\(|\)|\.|\s', '', g.person.telephone())
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
    batch_size = 1000
    success, fail = 0, 0
    g = Generic('en')
    r = random.Random()
    ts_base = int(time.mktime(datetime.date(2019, 1, 1).timetuple()))

    insert_stmt = '''
insert into casa_account (account_id, nick_name, prod_code, prod_name, currency, status, balance)
values (?, ?, ?, ?, ?, ?, ?)
    '''

    cur = sess.cursor(prepared=True)
    for i in range(n):
        account_id = unique_account_id(g)
        prod_code, prod_name = random_product(r)
        balance = r.uniform(1000.0, 5_000_000.0, 4)
        # ts = int(r.uniform(0, 15_724_800, 0))  # 3600 * 24 * 182
        # update_dt = datetime.datetime.fromtimestamp(ts_base + ts)
        # print(update_dt)
        try:
            val = (account_id, g.person.name(), prod_code, prod_name, 'THB', 0, balance)
            cur.execute(insert_stmt, val)

            if i % batch_size == 0 or i == n-1:
                sess.commit()
                success += batch_size
                print(f'commiting {batch_size} records, total records = {success}')

        except mysql.connector.Error as e:
            print(e)

    return success, fail


def populate_testdata(n, force_drop):
    sess = create_session()
    print_db_version(sess)
    # sys.exit(0)
    if force_drop:
        drop_and_recreate_table(sess)

    success, fail = gen_random_accounts(sess, n)
    with open('ids.txt', 'w') as f:
        f.writelines('%s\n' % id for id in all_account_ids)
    print(f'success={success}, fail={fail}, uniq_id={len(all_account_ids)}, accounts_in_db={accounts_in_db(sess)})')

    sess.disconnect()


def drop_and_recreate_table(sess):
    # drop existing table and UDT
    cur = sess.cursor()
    cur.execute('DROP TABLE IF EXISTS casa_account')
    # recreate
    cur.execute('''
create table casa_account
(
    account_id varchar(255) not null,
    nick_name  varchar(255) not null,
    prod_code  varchar(255) not null,
    prod_name  varchar(255) not null,
    currency   varchar(3)   not null,
    status     int          not null,
    balance    float        not null,
    constraint account_id   unique (account_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4
    ''')


def create_session():
    return mysql.connector.connect(**db_conf)


def print_db_version(sess):
    print(sess.get_server_version())


def accounts_in_db(sess):
    cur = sess.cursor()
    cur.execute('select count(*) from casa_account')
    return cur.rowcount


def init_argparse():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '-v', '--version', action='version',
        version=f'{parser.prog} version 1.0.0'
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
