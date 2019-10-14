let SessionLoad = 1
let s:so_save = &so | let s:siso_save = &siso | set so=0 siso=0
let v:this_session=expand("<sfile>:p")
silent only
cd ~/Sync/3DP/github/goduetapiclient/machine
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
set shortmess=aoO
badd +1 move.go
badd +2 state.go
badd +17 tool.go
badd +947 mm.json
badd +39 ~/Sync/3DP/github/goduetapiclient/cmd/test/test.go
argglobal
%argdel
$argadd channels.go
$argadd electronics.go
$argadd fan.go
$argadd heat.go
$argadd laser.go
$argadd machinemodel.go
$argadd messagebox.go
$argadd move.go
$argadd network.go
$argadd scanner.go
$argadd sensors.go
$argadd spindle.go
$argadd state.go
$argadd storage.go
$argadd tool.go
$argadd uservariable.go
edit move.go
set splitbelow splitright
set nosplitbelow
wincmd t
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
argglobal
if bufexists("move.go") | buffer move.go | else | edit move.go | endif
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let s:l = 1 - ((0 * winheight(0) + 34) / 69)
if s:l < 1 | let s:l = 1 | endif
exe s:l
normal! zt
1
normal! 0
lcd ~/Sync/3DP/github/goduetapiclient/machine
tabnext 1
if exists('s:wipebuf') && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20 winminheight=1 winminwidth=1 shortmess=filnxtToOFc
let s:sx = expand("<sfile>:p:r")."x.vim"
if file_readable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &so = s:so_save | let &siso = s:siso_save
let g:this_session = v:this_session
let g:this_obsession = v:this_session
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
