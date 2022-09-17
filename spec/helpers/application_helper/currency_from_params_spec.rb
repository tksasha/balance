# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#currency_from_params' do
    let(:params) { { currency: 'uah' } }

    its(:currency_from_params) { is_expected.to eq 'uah' }
  end
end
