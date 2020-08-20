import BaseHTTPServer
import SocketServer
import json
import time


HOST_NAME = '0.0.0.0'
PORT_NUMBER =8080

customer = json.dumps({
	'customer_id':'10001000',
	'name' : '10001000',
	'login_name' : '10001000'
})

class MyHandler(BaseHTTPServer.BaseHTTPRequestHandler):
    def do_GET(s):
        s.send_response(200)
        s.send_header('Content-type', 'application/json')
        s.end_headers()
        s.wfile.write(customer)

if __name__ == '__main__':
  httpd = BaseHTTPServer.HTTPServer((HOST_NAME, PORT_NUMBER), MyHandler)
  print(time.asctime(), "Server Starts - %s:%s" % (HOST_NAME, PORT_NUMBER))
  try:
      httpd.serve_forever()
  except KeyboardInterrupt:
      pass
  httpd.server_close()
