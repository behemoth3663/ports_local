PORTNAME=	tailcat
PORTVERSION=	0.5.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	netcat-like tool over Tailscale's data plane, no control plane
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tailscale

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=bradfitz GHT_PROJ=wintun-go GHT_PATH=golang.zx2c4.com/wintun )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-go4org:go4:ae6ca9944745:go4org_go4/vendor/go4.org/mem
#+go4org:mem:ae6ca9944745:go4org_go4/vendor/go4.org/mem
#-go4org:go4:fdeea329fbba:go4org_go4_1/vendor/go4.org/netipx
#+go4org:netipx:fdeea329fbba:go4org_go4_1/vendor/go4.org/netipx
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:)[^:]+:(.*/vendor/go4.org)/(mem|netipx),\1\3:\2/\3,}

#-wireguard:wireguard-go:v0.5.3:wireguard_wireguard_go/vendor/golang.zx2c4.com/wireguard/windows
#+WireGuard:wireguard-windows:v0.5.3:wireguard_wireguard_go/vendor/golang.zx2c4.com/wireguard/windows
GH_TUPLE:=	${GH_TUPLE:S|^wireguard:wireguard-go:|WireGuard:wireguard-windows:|1}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
