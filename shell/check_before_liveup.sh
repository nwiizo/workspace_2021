 mkdir bk_$(date "+%Y%m%d")
 cd bk_$(date "+%Y%m%d")
 ip a > before_ip.txt
 ip r > before_route.txt
 df -h > before_df.txt
 netstat -natp | grep LISTEN > before_netstat.txt
 ps awxu | awk '{for(i=11;i<NF;i++){printf("%s%s",$i,OFS=" ")}print $NF}' > before_ps_custom.txt
 iptables-save > before_iptables.txt
