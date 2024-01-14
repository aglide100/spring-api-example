package com.example.exception;

        import com.example.model.msg.MessageStatus;

        import java.util.Locale;


public class UnAuthorized extends MessageException {
    public UnAuthorized() {
        super(MessageStatus.RES_HTTP_UNAUTHORIZED);
    }

    public UnAuthorized(Locale locale) {
        super(MessageStatus.RES_HTTP_UNAUTHORIZED.localized(locale));
    }
}