import React, { FunctionComponent, useState } from 'react';
import { debounce } from 'lodash';
import { Input } from '@grafana/ui';
import { InlineFields } from '.';

import { INPUT_WIDTH, LABEL_WIDTH } from '../constants';

export interface Props {
  onChange: (alias: any) => void;
  value?: string;
}

export const AliasBy: FunctionComponent<Props> = ({ value = '', onChange }) => {
  const [alias, setAlias] = useState(value ?? '');

  const propagateOnChange = debounce(onChange, 1000);

  onChange = (e: any) => {
    setAlias(e.target.value);
    propagateOnChange(e.target.value);
  };

  return (
    <InlineFields label="Alias By" grow transparent labelWidth={LABEL_WIDTH}>
      <Input width={INPUT_WIDTH} type="text" value={alias} onChange={onChange} />
    </InlineFields>
  );
};
