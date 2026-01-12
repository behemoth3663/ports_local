PORTNAME=	rad
PORTVERSION=	0.6.27
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local lang

MAINTAINER=	none@FreeBSD.org
COMMENT=	Modern CLI scripts made easy
WWW=		https://amterp.github.io/rad/

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	amterp

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@test ! -f ${WRKSRC}/vendor/github.com/tree-sitter/go-tree-sitter/dup_unix.go || ${REINPLACE_CMD} 's/go:build/& freebsd ||/' ${WRKSRC}/vendor/github.com/tree-sitter/go-tree-sitter/dup_unix.go
	@if [ -d ${WRKSRC}/vendor ]; then \
		${MKDIR} ${WRKSRC}/vendor/github.com/charmbracelet/x/exp; \
		${RLN} ${WRKSRC_charmbracelet_x_strings}/exp/strings ${WRKSRC}/vendor/github.com/charmbracelet/x/exp; \
		${RLN} ${WRKSRC_charmbracelet_x_ansi}/ansi ${WRKSRC}/vendor/github.com/charmbracelet/x; \
		${RLN} ${WRKSRC_charmbracelet_x_cellbuf}/cellbuf ${WRKSRC}/vendor/github.com/charmbracelet/x; \
		${RLN} ${WRKSRC_charmbracelet_x_term}/term ${WRKSRC}/vendor/github.com/charmbracelet/x; \
	fi

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-charmbracelet:x:b4e2667e5ad5:charmbracelet_x_strings/vendor/github.com/charmbracelet/x/exp/strings
#+charmbracelet:x:b4e2667e5ad5:charmbracelet_x_strings
GH_TUPLE:=	${GH_TUPLE:S,/vendor/github.com/charmbracelet/x/exp/strings,,1}

#-charmbracelet:x:v0.0.13:charmbracelet_x_cellbuf/vendor/github.com/charmbracelet/x/cellbuf
#-charmbracelet:x:v0.10.1:charmbracelet_x_ansi/vendor/github.com/charmbracelet/x/ansi
#-charmbracelet:x:v0.2.1:charmbracelet_x_term/vendor/github.com/charmbracelet/x/term
#+charmbracelet:x:ansi/v0.10.1:charmbracelet_x_ansi \
#+charmbracelet:x:cellbuf/v0.0.13:charmbracelet_x_cellbuf \
#+charmbracelet:x:term/v0.2.1:charmbracelet_x_term \
#GH_TUPLE:=	${GH_TUPLE:C,(:charmbracelet_x_(ansi|cellbuf|term)).*,\1,}
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+).*/charmbracelet/x/(ansi|cellbuf|term),\1\3/\2,}

#-tree-sitter:go-tree-sitter:c949200:tree_sitter_go_tree_sitter/vendor/github.com/tree-sitter/go-tree-sitter
#+tree-sitter:go-tree-sitter:v0.25.0:tree_sitter_go_tree_sitter/vendor/github.com/tree-sitter/go-tree-sitter
GH_TUPLE:=	${GH_TUPLE:C,^(tree-sitter:go-tree-sitter:)[^:]+(.*),\1c949200\2,}

.include <bsd.port.mk>
