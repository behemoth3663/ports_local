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

GH_SHORT_HASH=	git ls-remote --tags https://github.com/${GH_ACCOUNT}/${PORTNAME}.git \
	| ${SED} -rn 's|^(.{7}).*/${DISTVERSIONPREFIX}${PORTVERSION:S/./\\&/g}$$|\1|p'

GO_BUILDFLAGS+=	-ldflags="-s \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildVersion=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildDate=${:!date +%F_%T!:S/_/ /}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildHash=${:!${GH_SHORT_HASH}!}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/cmd.BuildEnvironment='"

post-patch: .PHONY
.if ${.TARGETS:Mgomod-vendor}
	cd ${WRKSRC} && ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} generate ./...
.endif

.include <bsd.port.mk>
