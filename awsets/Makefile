# $FreeBSD$

PORTNAME=	awsets
PORTVERSION=	0.8.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for export all AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/../../LICENSE

USES=		go:modules

WRKSRC_SUBDIR=	cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	trek10inc

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-X 'main.version=${PORTVERSION}' -X 'main.commit=8d9d922' -X 'main.date=${:!date +%F_%T!:S/_/ /}'"

fetch-version: .PHONY get-version

post-patch: .PHONY
	${MKDIR} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}
	${FIND} ${WRKSRC}/../.. -maxdepth 1 \( \
		\( -type f -and -name '*.go' \) -or \
		\( -type d -and -not -name '.*' -and -not -name 'cmd' \) \
	\) -exec ${RLN} {} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME} \;

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
