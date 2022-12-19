import asyncio
from websockets import connect


async def hello(uri):
    async with connect(uri) as websocket:
        while True:
            x = await websocket.recv()
            print(x)


asyncio.run(hello("ws://localhost:1323/ws/"))
