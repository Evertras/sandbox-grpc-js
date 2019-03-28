import { messages } from './proto/echo.pb';
import { createEcho } from './proto/echo.twirp';

console.log('Hello from TS Land!');

const echoService = createEcho('http://localhost:8000');

async function echo(greeting: string) {
    const response = await echoService.sayHello( <messages.IEchoMessage> { greeting });

    console.log(response.greeting);
}

echo('hello');
