# Announce

Go server which plays audio generated from AWS Polly through the servers
3.5mm headphone jack. Intended for raspberry pi for use with IFTTT to have
messages announced over loudspeakers.

### Configuration

Must set AWS Access key, AWS Secret key, and port env variables
* `AWS_ACCESS_KEY`
* `AWS_SECRET_KEY`
* `PORT`

### Example uses

Announcing when leaving work

> Brian has left work and is on his way home


Announcing weather

> Good morning. Today's weather will be partly cloudy with a high of 52 and a low of 40


Announcing Twitter activity

> You just got a new twitter follower. @someuser
