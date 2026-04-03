PORTNAME=	rad
PORTVERSION=	0.12.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local lang

MAINTAINER=	none@FreeBSD.org
COMMENT=	Modern CLI scripts made easy
WWW=		https://amterp.github.io/rad/ https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	amterp

CGO_CFLAGS=	-I${WRKSRC}/vendor/github.com/tree-sitter/go-tree-sitter/include

GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-extldflags -static
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

M2T_OFFLINE=	true
.export	M2T_OFFLINE

PLIST_FILES=	bin/${PORTNAME}

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${MKDIR} ${WRKSRC}/vendor
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${TEST} ! -f ${WRKSRC}/vendor/github.com/tree-sitter/go-tree-sitter/dup_unix.go || ${REINPLACE_CMD} 's/go:build/& freebsd ||/' ${WRKSRC}/vendor/github.com/tree-sitter/go-tree-sitter/dup_unix.go
	@${MKDIR} ${WRKSRC}/vendor/github.com/amterp/tree-sitter-rad
	@if [ -n '${WRKSRC_amterp_tree_sitter_rad}' ]; then \
		${RLN} ${WRKSRC_amterp_tree_sitter_rad}/bindings ${WRKSRC}/vendor/github.com/amterp/tree-sitter-rad; \
		${RLN} ${WRKSRC_amterp_tree_sitter_rad}/src ${WRKSRC}/vendor/github.com/amterp/tree-sitter-rad; \
	fi

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-tree-sitter:go-tree-sitter:v0.25.0:tree_sitter_go_tree_sitter/vendor/github.com/tree-sitter/go-tree-sitter
#+tree-sitter:go-tree-sitter:c949200:tree_sitter_go_tree_sitter/vendor/github.com/tree-sitter/go-tree-sitter
GH_TUPLE:=	${GH_TUPLE:C,^(tree-sitter:go-tree-sitter:)[^:]+(.*),\1c949200\2,1}

#-amterp:tree-sitter-rad:v0.10.0:amterp_tree_sitter_rad/vendor/github.com/amterp/tree-sitter-rad
#+amterp:tree-sitter-rad:v0.10.0:amterp_tree_sitter_rad
GH_TUPLE:=	${GH_TUPLE:S,/vendor/github.com/amterp/tree-sitter-rad,,1}

.include <bsd.port.mk>
