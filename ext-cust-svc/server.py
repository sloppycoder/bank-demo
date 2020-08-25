from aiohttp import web


async def handle(request):
    cid = request.match_info.get('id')
    # print(request.headers)
    return web.json_response(gen_customer(cid))


def gen_customer(cust_id):
    return {
        'customer_id': cust_id,
        'name': 'dummy',
        'login_name': cust_id
    }


app = web.Application()
app.router.add_get('/customers/{id}', handle)

web.run_app(app)
