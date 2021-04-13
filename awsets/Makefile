# $FreeBSD$

PORTNAME=	awsets
PORTVERSION=	0.8.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for export all AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/../../LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

WRKSRC_SUBDIR=	cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	trek10inc

PLIST_FILES=	bin/${PORTNAME}

#8d9d922d9dc592bedd5ac77b89c71a5f5b1b7e1c	refs/tags/v0.8.0
GH_SHORT_HASH=	git ls-remote --tags https://github.com/${GH_ACCOUNT}/${PORTNAME}.git \
	| ${SED} -rn 's|^(.{7}).*/${DISTVERSIONPREFIX}${PORTVERSION:S/./\\&/g}$$|\1|p'

GO_BUILDFLAGS+=	-ldflags=" \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.commit=${:!${GH_SHORT_HASH}!}' \
	-X 'main.date=${:!date +%F_%T!:S/_/ /}'"

fetch-version: .PHONY get-version

post-patch: .PHONY
	${MKDIR} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}
	${FIND} ${WRKSRC}/../.. -maxdepth 1 \( \
		\( -type f -and -name '*.go' \) -or \
		\( -type d -and -not -name '.*' -and -not -name 'cmd' \) \
	\) -exec ${RLN} {} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME} \;

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
