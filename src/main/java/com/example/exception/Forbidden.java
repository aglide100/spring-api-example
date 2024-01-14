package com.example.exception;

import com.example.model.msg.MessageStatus;
import java.util.Locale;

public class Forbidden extends Common {
    public Forbidden() {
        super(MessageStatus.RES_HTTP_FORBIDDEN);
    }

    public Forbidden(Locale locale) {
        super(MessageStatus.RES_HTTP_FORBIDDEN.localized(locale));
    }
}