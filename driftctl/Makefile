#

PORTNAME=	driftctl
PORTVERSION=	0.19.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Detect, track and alert on infrastructure drift

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules

GO_MODULE=	github.com/snyk/${PORTNAME}
GH_ACCOUNT=	${GO_MODULE:H:T}
GH_PROJECT=	${GO_MODULE:T}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X '${GO_MODULE}/pkg/version.version=${PORTVERSION}' \
	-X '${GO_MODULE}/build.env=release' \
	-X '${GO_MODULE}/build.enableUsageReporting=false'"

fetch-version: .PHONY get-version

.include <bsd.port.mk>
