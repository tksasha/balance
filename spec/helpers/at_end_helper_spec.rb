# frozen_string_literal: true

RSpec.describe AtEndHelper, type: :helper do
  subject { helper }

  describe '#at_end' do
    let(:params) { { currency: 'uah' } }

    before { allow(subject).to receive(:params).and_return(params) }

    before { allow(CalculateAtEndService).to receive(:call).with('uah').and_return(21.49) }

    its(:at_end) { is_expected.to eq 21.49 }
  end
end
