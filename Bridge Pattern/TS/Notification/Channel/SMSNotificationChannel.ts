import {NotificationChannel} from './NotificationChannel'

export class SMSNotificationChannel implements NotificationChannel {
    sendNotification(recipient: string, message: string): void {
        console.log(`LOG: Messages has sent successfully to ${recipient} via sms. \nMessage: ${message} \n`)
    };
}