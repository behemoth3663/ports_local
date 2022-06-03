#

PORTNAME=	ratchet
PORTVERSION=	0.2.3
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	A tool for securing CI/CD workflows with version pinning

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sethvargo

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

#GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
#	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
#	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}'"

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps
#doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

post-patch:
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*Version[[:space:]]*=[[:space:]]*).*|\1"${PORTVERSION}"|' \
		-e 's|^([[:space:]]*Commit[[:space:]]*=[[:space:]]*).*|\1"${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}"|' \
	${WRKSRC}/internal/version/version.go
 
.include <bsd.port.mk>
