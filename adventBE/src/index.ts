import {App} from './application';
import {routes} from './routes/router';
import * as dotenv from 'dotenv';

dotenv.config();

const port = 4000;

const app = new App(
port, [routes]);

app.listen();


