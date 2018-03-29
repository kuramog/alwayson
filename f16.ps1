Add-Type -AssemblyName System.Windows.Forms

$job = Start-Job {
    while(1){
        [System.Windows.Forms.SendKeys]::SendWait("{F16}")
        Start-Sleep -Seconds 59
    }
}

[System.Windows.Forms.MessageBox]::Show("exit?","alwayson")

$job|Stop-Job
