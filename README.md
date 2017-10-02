# ulmf

Ugh, Let Me Focus! (a helper to shut up Slack)

In order to mark yourself as busy or away for focusing, Slack requires you to do three things:

1. `/away`
2. `/status :computer: busy`
3. `/dnd 120`

And then, at the end you must do the inverse.

This command does all three, via the API.

# Usage

```
$ ulmf -t xoxp-1234 -d 120 -e ":computer:" -m "focus time"
```

The slack token can also be specified as an environment variable:

```
$ SLACK_TOKEN=xoxp-1234 ulmf -d 120 -e ":computer:" -m "focus time"
```

And to undo your away:

```
$ ulmf -t xoxp-1234 -u
```
