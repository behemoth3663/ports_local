PORTNAME=	iniq
DISTVERSIONPREFIX=	v
DISTVERSION=	0.4.0
PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Simple INI file reader for command line
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		gmake perl5
USE_PERL5=	build

USE_GITHUB=	yes
GH_ACCOUNT=	jcrd

MAKE_ARGS=	PREFIX=${PREFIX}

PLIST_FILES=	bin/${PORTNAME} \
		share/man/man1/${PORTNAME}.1.gz

OPTIONS_DEFINE=	STATIC TEST
OPTIONS_DEFAULT=STATIC TEST

STATIC_MAKE_ARGS=	LDFLAGS=-static

post-build-TEST-on: test-message .PHONY
	${MAKE_CMD} -C ${WRKSRC} test

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
