#!/usr/bin/env python
import smtplib
import sys
content = 'Mensaje de notificacion '+sys.argv[1]
asunto = ''
mail = smtplib.SMTP('smtp-pulse.com',2525)
para = 'cerdolot13@gmail.com'
desde = 'soporte@awaresystems.cl'

header = 'To:' + para + '\n' + 'From: ' + desde + '\n' + 'Subject:testing \n'
msg = header + '\n Este correo sera el futuro del QoS para pruebas Aware Systems\n\n'

mail.ehlo()

mail.starttls()

mail.login('sebacav@gmail.com', 'DH7T4RHsKe')

mail.sendmail(desde,para,msg)

mail.close()
