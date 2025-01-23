export interface NotificationChannel {
    sendNotification(recipient: string, message: string): void;
}