import {SMSNotificationChannel} from "./Notification/Channel/SMSNotificationChannel";
import {EmailNotificationChannel} from "./Notification/Channel/EmailNotificationChannel";
import {PatientVerifiedNotification} from "./Notification/PatientVerifiedNotification";
import {NotificationChannel} from "./Notification/Channel/NotificationChannel";
import {PatientRejectedNotification} from "./Notification/PatientRejectedNotification";


const activeChannel: string = "sms";

function main() {
    let channel: NotificationChannel | null = null;


    switch (activeChannel) {
        case "sms":
            channel = new SMSNotificationChannel();
            break;
        case "email":
            channel = new EmailNotificationChannel();
            break;
        default:
            return
    }

    new PatientVerifiedNotification(channel).send("jhondoe@gmail.com", {
        patientId: 1
    })

    new PatientRejectedNotification(channel).send("alex@yahoo.com", {
        patientId: 2
    })
}


main();