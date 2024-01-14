package com.example.model.msg;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.support.MessageSourceAccessor;
import org.springframework.stereotype.Component;

import java.util.Locale;

@Component
public class
common {

        private static MessageSourceAccessor i18n;

        @Autowired
        public void messageSourceAcc(MessageSourceAccessor msAcc) {
            this.i18n = msAcc;
        }

        public static String getMessageDetail(String code) {
            return i18n.getMessage(code);
        }

        public static String getMessageDetail(String code, Object[] objs) {
            return i18n.getMessage(code, objs);
        }

        public static String getMessageDetail(String code, Object[] objs, Locale locale) {
            return i18n.getMessage(code, objs, locale);
        }
}
