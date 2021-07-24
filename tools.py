# ------------------------------------------------------------------------------
# Some tools to deal with facebook archive json.
# ------------------------------------------------------------------------------

import json
from datetime import datetime

def calc_total_messages():
    total_msg = 0
    mList = []

    # Generate file name
    for i in range(1, 21):
        file_name = 'message_'
        extension = '.json'
        full_name = file_name + str(i) + extension
        mList.append(full_name)

    # Read messages
    for i in mList:
        with open(i) as f:
            data = json.load(f)
            a = data['messages']
            total_msg += len(a)
            print(len(a))

    print(total_msg)

def encode_to_human_raedable():
    data = r'"\u00f0\u009f\u0094\u0091"'
    return json.loads(data).encode('latin1').decode('utf8')


# Print message using followin format
# 1970-01-01 00:00:00 <Sender> Message
with open('message_20.json') as f:
    data = json.load(f)

    for i in data['messages']:
        # Timestamp
        ts = i['timestamp_ms'] / 1000
        print(datetime.fromtimestamp(ts).strftime('%Y-%m-%d %H:%M:%S'), end=' ')

        # Sender
        sender_name = i['sender_name'].encode('latin1').decode('utf8')
        print('<%s>' % sender_name, end=' ')

        # Message
        if 'content' in i:
            print(i['content'].encode('latin1').decode('utf8'))
        elif 'photos' in i:
            print(i['photos'][0]['creation_timestamp'], i['photos'][0]['uri'])
        elif 'sticker' in i:
            print(i['sticker'])
