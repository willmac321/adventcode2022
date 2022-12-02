"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.App = void 0;
const express = require("express");
class App {
    /**
     * @param port Port Application listens on
     * @param middleware Array of middleware to be applied to app
     * @param routes Array of express.Router objects for application routes
     * @param apiPath Base path for this api that will be prepended to all routes
     * @param staticPath path to folder for public files express will make available
     */
    constructor(port, routes, apiPath = '/api', staticPath = "public") {
        this.port = port;
        this.apiPath = apiPath;
        this.staticPath = staticPath;
        //* Create a new express app
        this.app = express();
        //* Method calls `this.app.use()` for each router, prepending `this.apiPath` to each router
        this.routes(routes);
        //* Method calls `this.app.use(express.static(path))` to enable public access to static files
        this.assets(this.staticPath);
    }
    /**
         * Attaches route objects to app, appending routes to `apiPath`
         * @param routes Array of router objects to be attached to the app
         */
    routes(routes) {
        routes.forEach((r) => {
            this.app.use(`${this.apiPath}`, r);
        });
    }
    /**
     * Enable express to serve up static assets
     */
    assets(path) {
        this.app.use(express.static(path));
    }
    /**
     * Start the Express app
     */
    listen() {
        this.app.listen(this.port, () => {
            console.log(`⚡️[server]: Server is running at https://localhost:${this.port}`);
        });
    }
}
exports.App = App;
