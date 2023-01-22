
.PHONY: fonts
fonts:
	mkdir -p ~/.local/Github

	# Material Design Icons
	git clone --depth 1 https://github.com/google/material-design-icons $$HOME/.local/Github/material-design-icons
	ln -s ~/.local/Github/material-design-icons/font ~/.local/share/fonts/material-design-icons-font

	# Community Fork
	git clone --depth 1 https://github.com/Templarian/MaterialDesign-Webfont $$HOME/.local/Github/MaterialDesign-Webfont
	ln -s ~/.local/Github/MaterialDesign-Webfont/fonts ~/.local/share/fonts/materialdesign-webfont

	# FontAwesome
	git clone --depth 1 https://github.com/FortAwesome/Font-Awesome $$HOME/.local/Github/Font-Awesome
	ln -s ~/.local/Github/Font-Awesome/otfs ~/.local/share/fonts/font-awesome

	# Typicons
	git clone --depth 1 https://github.com/stephenhutchings/typicons.font $$HOME/.local/Github/typicons
	ln -s ~/.local/Github/typicons.font/src/font ~/.local/share/fonts/typicons-font

.PHONY: build
build:
	@DOCKER_BUILDKIT=1 docker image build \
		--network host \
		--output type=local,dest=bin \
		.
