# frozen_string_literal: true

RSpec.describe AtEndHelper, type: :helper do
  subject { helper }

  describe '#at_end' do
    let(:params) { { currency: 'uah' } }

    before do
      allow(subject).to receive(:params).and_return(params)

      allow(CalculateAtEndService).to receive(:call).with('uah').and_return(21.49)
    end

    its(:at_end) { is_expected.to eq 21.49 }
  end
end
