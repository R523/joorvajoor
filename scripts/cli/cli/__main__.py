import socketserver
import argparse

from python_mpv_jsonipc import MPV

parser = argparse.ArgumentParser(description='remove controlled player')
parser.add_argument('movie', metavar="movie", type=str, nargs='?', help='movie you want to play', default='/Users/parham/Downloads/SampleVideo_1280x720_30mb.mp4')
args = parser.parse_args()

mpv = MPV()
mpv.play(args.movie)

class MPVHandler(socketserver.StreamRequestHandler):
    def handle(self) -> None:
        cmd = self.rfile.readline().strip()
        print(f'received command is {cmd}')
        if cmd == b'play':
            print(mpv.command('set_property', 'pause', False))
        elif cmd == b'pause':
            print(mpv.command('set_property', 'pause', True))
        else:
            print('command not found')

server = socketserver.TCPServer(("0.0.0.0", 1378), MPVHandler)

try:
    server.serve_forever()
except KeyboardInterrupt:
    server.server_close()
    mpv.terminate()
