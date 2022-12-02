import {App} from './application';
import {routes} from './routes/router';

const port = 4000;

const app = new App(
port, [routes]);

app.listen();


