def deep_update(src, update):
    '''
    type(src) == dict
    type(update) == dict
    use dict `update` to recursively update dict `src`
    '''
    if not isinstance(src, dict) or not isinstance(update, dict):
        print 'Invalid parameters!'
        return

    for k, v in update.iteritems():
        if k in src and v != src[k]:
            if isinstance(src[k], dict) and isinstance(v, dict):
                deep_update(src[k], v)
            else:
                src[k] = v
        else:
            src[k] = v
