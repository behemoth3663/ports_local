PORTNAME=	terraform-docs
PORTVERSION=	0.19.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	3
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules,1.23

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-docs

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-		google:go-genproto:3a041ad873d4:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+		google:go-genproto:3a041ad873d4:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||}

.include <bsd.port.mk>
