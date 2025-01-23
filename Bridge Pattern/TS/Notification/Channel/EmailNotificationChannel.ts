import {NotificationChannel} from './NotificationChannel'

export class EmailNotificationChannel implements NotificationChannel {
    sendNotification(recipient: string, message: string): void {
        console.log(`Messages has sent successfully to ${recipient} via email. \n message: ${message}`)
    };
}