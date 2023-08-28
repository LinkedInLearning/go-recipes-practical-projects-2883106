#!/usr/bin/perl
# Dummy HTTP server. Writes it's PID to httpd.pid

my $pid_file = "httpd.pid";
open(my $fh, ">", "httpd.pid") || die "can't open $pid_file";
print $fh $$;
close $fh;
print "Server running (pid=$$)";
sleep(10000);
