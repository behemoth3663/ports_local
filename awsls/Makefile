# $FreeBSD$

PORTNAME=	awsls
PORTVERSION=	0.10.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	List command for AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.md

RUN_DEPENDS=	git:devel/git

#with go:modules doesn't build
USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	jckuester

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

GH_SHORT_HASH=	git ls-remote --tags https://github.com/${GH_ACCOUNT}/${PORTNAME}.git \
	| ${SED} -rn 's|^(.{7}).*/${DISTVERSIONPREFIX}${PORTVERSION:S/./\\&/g}$$|\1|p'

#doesn't work
#GO_BUILDFLAGS+=-ldflags="\
#	-X 'internal.version=${PORTVERSION}' \
#	-X 'internal.commit=${:!${GH_SHORT_HASH}!}' \
#	-X 'internal.date=${:!date +%F_%T!:S/_/ /}'"
do-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*version[[:space:]]*=[[:space:]]*")[^"]*(".*)|\1${PORTVERSION}\2|' \
		-e 's|^([[:space:]]*commit[[:space:]]*=[[:space:]]*")[^"]*(".*)|\1${:!${GH_SHORT_HASH}!}\2|' \
		-e 's|^([[:space:]]*date[[:space:]]*=[[:space:]]*")[^"]*(".*)|\1${:!date +%F_%T!:S/_/ /}\2|' \
	${WRKSRC}/internal/version.go

do-build: .PHONY
	pwd="$${PWD}" && trap "cd \"$${pwd}\"" INT EXIT && cd ${WRKSRC} && \
	${GO_CMD} build ${GO_BUILDFLAGS} -o ${GO_WRKDIR_BIN}/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
