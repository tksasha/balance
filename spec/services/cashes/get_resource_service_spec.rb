# frozen_string_literal: true

RSpec.describe Cashes::GetResourceService do
  subject { described_class.new params }

  let(:params) { { id: 27 } }

  describe '#call' do
    before { allow(Cash).to receive(:find).with(27).and_return(:cash) }

    its(:call) { should eq :cash }
  end
end
