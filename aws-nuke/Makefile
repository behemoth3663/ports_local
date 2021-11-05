#

PORTNAME=	aws-nuke
PORTVERSION=	2.16.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Remove all resources from an AWS account

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	rebuy-de

PLIST_FILES=	bin/${PORTNAME}

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps
#doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildVersion=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildDate=${%FT%TZ:L:gmtime}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildHash=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildEnvironment='"

post-patch: .PHONY
.if ${.TARGETS:Mgomod-vendor}
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} generate ./...
.endif

.include <bsd.port.mk>
