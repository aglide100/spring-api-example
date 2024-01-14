package com.example.model.msg;

import lombok.Getter;

import java.util.Locale;

@Getter
public enum MessageStatus {
    RES_HTTP_OK(200, "api.response.msg.200"),
    RES_HTTP_ACCEPTED(202, "api.response.msg.202"),
    RES_HTTP_BAD_REQUEST(400, "api.response.msg.400"),
    RES_HTTP_UNAUTHORIZED(401, "api.response.msg.401"),
    RES_HTTP_FORBIDDEN(403, "api.response.msg.403"),
    RES_HTTP_NOT_FOUND(404, "api.response.msg.404"),
    RES_HTTP_PRECONDITION_FAILED(412, "api.response.msg.412"),
    RES_HTTP_SERVER_ERROR(500, "api.response.msg.500");

    private final int status;
    private final String code;
    private Object[] objects;
    private String detail;

    MessageStatus(int status, String messageCode) {
        this.status = status;
        this.code = messageCode;
        this.detail = common.getMessageDetail(messageCode);
    }

    public MessageStatus localized(Locale locale) {
        this.detail = common.getMessageDetail(this.code, this.objects, locale);
        return this;
    }
}
