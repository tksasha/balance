# frozen_string_literal: true

RSpec.describe Cashes::InitializeResourceService do
  subject { described_class.new }

  its(:cash) { should be_a Cash }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    its(:call) { should be_success }

    its('call.object') { should eq cash }
  end
end
