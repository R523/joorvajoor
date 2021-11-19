import socketserver

from python_mpv_jsonipc import MPV

mpv = MPV()

class MPVHandler(socketserver.StreamRequestHandler):
    def handle(self) -> None:
        cmd = self.rfile.readline()
        if cmd == 'play':
            mpv.command('set_property', 'pause', False)
        elif cmd == 'pause':
            mpv.command('set_property', 'pause', True)
        else:
            print('command not found')

with socketserver.TCPServer(("0.0.0.0", 1378), MPVHandler) as server:
    server.serve_forever()
