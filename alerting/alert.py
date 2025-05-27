import smtplib
from email.mime.text import MIMEText

threshold = 75
current_cpu = 82

if current_cpu > threshold:
    msg = MIMEText(f"CPU usage is high: {current_cpu}%")
    msg['Subject'] = 'High CPU Alert'
    msg['From'] = 'alert@example.com'
    msg['To'] = 'admin@example.com'

    with smtplib.SMTP('smtp.mailtrap.io', 2525) as server:
        server.login('user', 'password')
        server.sendmail(msg['From'], [msg['To']], msg.as_string())
