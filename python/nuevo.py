import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.header import Header
from email.utils import formataddr

msg = MIMEMultipart('alternative')
msg['Subject'] = 'Prueba de Notificacion de QoS'
msg['From'] = formataddr((str(Header('Aware Systems', 'utf-8')), 'soporte@awaresystems.cl'))
msg['To'] = 'web-uq0ym@mail-tester.com'
print msg['From']

html = "Prueba de notificacion de mail via Python"

# Record the MIME types of text/html.
msg.attach(MIMEText(html, 'html'))

# Send the message via local SMTP server.
s = smtplib.SMTP('awaresystems.cl',587)
s.ehlo()

s.starttls()

s.login('soporte@awaresystems.cl', 'lala123')

# sendmail function takes 3 arguments: sender's address, recipient's address
# and message to send - here it is sent as one string.
s.sendmail('soporte@awaresystems.cl', 'sebacav@gmail.com', msg.as_string())
s.quit()