import { Application } from "express";
import express = require("express");

export class App {
    public app: Application;

    /**
     * @param port Port Application listens on
     * @param middleware Array of middleware to be applied to app 
     * @param routes Array of express.Router objects for application routes
     * @param apiPath Base path for this api that will be prepended to all routes
     * @param staticPath path to folder for public files express will make available
     */
    constructor(
        readonly port: number,
        routes: express.Router[],
        readonly apiPath: string = '/api',
        readonly staticPath: string = "public"
    ) {
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
    private routes(routes: express.Router[]) :void {
        routes.forEach((r) => {
            this.app.use(`${this.apiPath}`, r);
        });
    }

    /**
     * Enable express to serve up static assets
     */
    private assets(path: string):void {
        this.app.use(express.static(path));
    }

    /**
     * Start the Express app
     */
    public listen():void {
        this.app.listen(this.port, () => {
            console.log(`⚡️[server]: Server is running at https://localhost:${this.port}`);
        });
    }

}
