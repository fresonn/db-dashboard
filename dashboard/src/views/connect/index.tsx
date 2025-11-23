import { useState } from 'react'
import { Input } from '@/components/ui/input'
import styles from './index.module.css'
import { Button } from '@/components/ui/button'
import postgresLogo from '@/assets/postgresql-logo.svg?url'
import postgresLogoDark from '@/assets/postgresql-logo-black.svg?url'
import { ThemeToggle } from '@/components/ui/theme-toggle'
import { useTheme } from '@/hooks/use-theme'
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput
} from '@/components/ui/shadcn/input-group'
import { Database, EthernetPort, Info, KeyRound, Server, UserRound } from 'lucide-react'
import { Label } from '@/components/ui/label'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/shadcn/popover'
import { TypographyH1, TypographySmall } from '@/components/ui/typography'

export function ConnectView() {
  const [theme] = useTheme()

  const [loading, setLoading] = useState(false)

  return (
    <main className="grid h-screen grid-cols-[2fr_1fr]">
      <section className="relative flex items-center justify-center dark:bg-neutral-900">
        <div className="absolute top-8 right-10">
          <ThemeToggle />
        </div>
        <div className="w-full max-w-[700px] px-12 py-10">
          <div className="flex items-center justify-center">
            <img
              src={theme === 'light' ? postgresLogoDark : postgresLogo}
              className="mr-2 size-16"
              alt="postgresql logo"
            />
            <TypographyH1>Simple Dashboard</TypographyH1>
          </div>
          <div className="mb-10 pt-10">
            <ul className="*:mb-4">
              <li>
                <Input label="Connection name" placeholder="My VPS postgres..." />
              </li>
              <li className="flex">
                <div className="mr-8">
                  <Label htmlFor="host-input">Host</Label>
                  <InputGroup>
                    <InputGroupInput id="host-input" placeholder="127.0.0.1" />
                    <InputGroupAddon>
                      <Server />
                    </InputGroupAddon>
                  </InputGroup>
                </div>
                <div>
                  <Label htmlFor="port-input">Port</Label>
                  <InputGroup>
                    <InputGroupInput id="port-input" type="number" placeholder="5432" />
                    <InputGroupAddon>
                      <EthernetPort />
                    </InputGroupAddon>
                  </InputGroup>
                </div>
              </li>
              <li>
                <Label htmlFor="user-input">User</Label>
                <InputGroup>
                  <InputGroupInput id="user-input" placeholder="postgres" />
                  <InputGroupAddon>
                    <UserRound />
                  </InputGroupAddon>
                </InputGroup>
              </li>
              <li>
                <Label htmlFor="password-input">Password</Label>
                <InputGroup>
                  <InputGroupInput id="password-input" type="password" />
                  <InputGroupAddon>
                    <KeyRound />
                  </InputGroupAddon>
                </InputGroup>
              </li>
              <li>
                <Label htmlFor="database-input">Database (optional)</Label>
                <InputGroup>
                  <InputGroupAddon>
                    <Database />
                  </InputGroupAddon>
                  <InputGroupInput id="database-input" placeholder="postgres" className="pl-2!" />
                  <InputGroupAddon align="inline-end">
                    <Popover>
                      <PopoverTrigger asChild>
                        <InputGroupButton className="cursor-pointer rounded-full" size="icon-xs">
                          <Info />
                        </InputGroupButton>
                      </PopoverTrigger>
                      <PopoverContent className="text-sm leading-normal">
                        By default it tries to connect to{' '}
                        <TypographySmall>postgres</TypographySmall> database
                      </PopoverContent>
                    </Popover>
                  </InputGroupAddon>
                </InputGroup>
              </li>
            </ul>
          </div>
          <Button loading={loading} onClick={() => setLoading((prev) => !prev)} fullWidth size="lg">
            Connect
          </Button>
        </div>
      </section>
      <section className={'flex items-center justify-center bg-white/90'}>
        <div className={styles.view} />
      </section>
    </main>
  )
}
