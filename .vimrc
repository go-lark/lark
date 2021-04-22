augroup LOCAL_SETUP
  autocmd FileType go lua require('go').setup{test_env = {GO_LARK_TEST_MODE = 'local'}}
augroup END
