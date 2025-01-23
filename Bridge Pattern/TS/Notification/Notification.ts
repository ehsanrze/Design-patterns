import {NotificationChannel} from './Channel/NotificationChannel'


export abstract class Notification {
    protected implementor: NotificationChannel;

    constructor(implementor: NotificationChannel) {
        this.implementor = implementor;
    }

    abstract send(recipient: string, inputs: any): void;
}