# $FreeBSD$

PORTNAME=	yc
PORTVERSION=	0.63.0
CATEGORIES=	sysutils linux
MASTER_SITES=   https://storage.yandexcloud.net/yandexcloud-yc/release/${DISTVERSION}/linux/${ARCH:S/i//}/
PKGNAMEPREFIX=	linux-
DISTNAME=	${PORTNAME}
EXTRACT_SUFX=	# none

MAINTAINER=	none@freebsd.org
COMMENT=        Yandex.Cloud command-line interface (CLI)

LICENSE=	PD

ONLY_FOR_ARCHS=	amd64 i386
ONLY_FOR_ARCHS_REASON=	Upstream does not provide binaries for other architerctures

USES=		linux

EXTRACT_CMD=	${CP}
EXTRACT_BEFORE_ARGS=
EXTRACT_AFTER_ARGS=	${WRKSRC}

NO_BUILD=	yes
NO_WRKSUBDIR=	yes

PLIST_FILES=	bin/${DISTNAME}

#fetch-version: .PHONY get-version
#get-version: .PHONY .SILENT
#	fetch -qo- https://storage.yandexcloud.net/yandexcloud-yc/release/stable

fetch-version: .PHONY .SILENT
	fetch -qo- https://storage.yandexcloud.net/yandexcloud-yc/release/stable

do-install: .PHONY .SILENT
.for f in ${PLIST_FILES}
	${BRANDELF} -t Linux ${WRKSRC}/${f:T}
	${INSTALL_PROGRAM} ${WRKSRC}/${f:T} ${STAGEDIR}${PREFIX}/${f:H:T}
.endfor

.include <bsd.port.mk>
