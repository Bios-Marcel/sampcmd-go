#include <Windows.h>
#include <string.h>
#include <stdio.h>

int _LaunchSAMPSetWorkingDir(wchar_t* working_dir, wchar_t* args) {
    int ret = 0;
    HMODULE module_handle = GetModuleHandleW(L"kernel32.dll");
    FARPROC load_library_w_proc;
    LPVOID ptr; 
    PROCESS_INFORMATION process_info;
    STARTUPINFOW startup_info;
    HANDLE rt;

    wchar_t samp_dll[FILENAME_MAX + 1];
    wchar_t gta_sa_exe[FILENAME_MAX + 1];

    memset(&process_info, 0, sizeof(PROCESS_INFORMATION));
    memset(&startup_info, 0, sizeof(STARTUPINFOW));

    int sz = wcslen(working_dir);
    memcpy(samp_dll, working_dir, sizeof(wchar_t) * sz);
    wcscat(samp_dll, L"\\samp.dll");

    memcpy(gta_sa_exe, working_dir, sizeof(wchar_t) * sz);
    wcscat(gta_sa_exe, L"\\gta_sa.exe");

    if (module_handle)
    {
        load_library_w_proc = GetProcAddress(module_handle, "LoadLibraryW");
        if (load_library_w_proc)
        {
            if (CreateProcessW(gta_sa_exe, args, NULL, NULL, FALSE, DETACHED_PROCESS | CREATE_SUSPENDED, NULL, NULL, &startup_info, &process_info))
            {
                ptr = VirtualAllocEx(process_info.hProcess, NULL, (wcslen(samp_dll) + 1) * sizeof(wchar_t), MEM_RESERVE | MEM_COMMIT, PAGE_EXECUTE_READWRITE);
                if (ptr)
                {
                    if (WriteProcessMemory(process_info.hProcess, ptr, samp_dll, (wcslen(samp_dll) + 1) * sizeof(wchar_t), NULL))
                    {
                        rt = CreateRemoteThread(process_info.hProcess, NULL, 0, (LPTHREAD_START_ROUTINE)load_library_w_proc, ptr, CREATE_SUSPENDED, NULL);
                        if (rt)
                        {
                            ResumeThread(rt);
                            WaitForSingleObject(rt, INFINITE);
                        }
                        else
                        {
                            ret = GetLastError();
                        }
                    }
                    else
                    {
                        ret = GetLastError();
                    }
                    VirtualFreeEx(process_info.hProcess, ptr, 0, MEM_RELEASE);
                }
                else
                {
                    ret = GetLastError();
                }
                ResumeThread(process_info.hThread);
                CloseHandle(process_info.hProcess);
            }
            else
            {
                ret = GetLastError();
            }
        }
        else
        {
            ret = GetLastError();
        }
    }
    else
    {
        ret = GetLastError();
    }

    return ret;
}

int _LaunchSAMP(wchar_t* args) {
    wchar_t cur_dir[FILENAME_MAX + 1];
    memset(cur_dir, 0, sizeof(wchar_t) * (FILENAME_MAX + 1));
    GetCurrentDirectoryW(FILENAME_MAX, cur_dir);

    return _LaunchSAMPSetWorkingDir((wchar_t * )(cur_dir[0]), args);
}
