PORTNAME=	oauth2-proxy
PORTVERSION=	7.4.0
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	www

MAINTAINER=	freebsd@rheinwolf.de
COMMENT=	OAuth 2.0 proxy server
WWW=		https://github.com/oauth2-proxy/oauth2-proxy

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_RC_SUBR=	oauth2-proxy

USE_GITHUB=	yes

USER=		www
GROUP=		www

VAR_DIR?=	/var
LOG_DIR?=	${VAR_DIR}/log/${PORTNAME}
RUN_DIR?=	${VAR_DIR}/run/${PORTNAME}

SUB_LIST+=	GROUP=${GROUP} \
		LOG_DIR=${LOG_DIR} \
		PORTNAME=${PORTNAME} \
		PREFIX=${PREFIX} \
		RUN_DIR=${RUN_DIR} \
		USER=${USER}

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${LOG_DIR}
	${MKDIR} ${STAGEDIR}${RUN_DIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/${PORTNAME}.cfg.example ${STAGEDIR}${PREFIX}/etc/${PORTNAME}.cfg.sample

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
