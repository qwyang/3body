#!/usr/bin/env python
# -*- coding: UTF-8 -*-

from email.mime.multipart import MIMEMultipart
from email.mime.base import MIMEBase
from email.mime.text import MIMEText
from email.utils import COMMASPACE, formatdate
from email import encoders
import os
import smtplib
import argparse


def send_mail(server, fro, to, subject, text, files=[]):
    assert type(server) == dict
    assert type(to) == list
    assert type(files) == list

    msg = MIMEMultipart()
    msg['From'] = fro
    msg['Subject'] = subject
    msg['To'] = COMMASPACE.join(to)  # COMMASPACE==', '
    msg['Date'] = formatdate(localtime=True)
    msg.attach(MIMEText(text))

    for file in files:
        part = MIMEBase('application', 'octet-stream')  # 'octet-stream': binary data
        part.set_payload(open(file, 'rb'.read()))
        encoders.encode_base64(part)
        part.add_header('Content-Disposition', 'attachment; filename="%s"' % os.path.basename(file))
        msg.attach(part)

    smtp = smtplib.SMTP(server['name'])
    smtp.starttls()
    smtp.login(server['user'], server['passwd'])
    smtp.sendmail(fro, to, msg.as_string())
    smtp.close()


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--user', required=True, help="qq email user")
    parser.add_argument('--passwd', required=True, help="qq email passwd")
    args = parser.parse_args()
    server = {'name': "smtp.qq.com", 'user': args.user, 'passwd': args.passwd}
    from_who = server['user'] + "@qq.com"
    to_list = ["yangqunwei@huawei.com"]
    send_mail(server, from_who, to_list, "test", "hello")
