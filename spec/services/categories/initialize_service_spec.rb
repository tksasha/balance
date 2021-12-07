# frozen_string_literal: true

RSpec.describe Categories::InitializeService do
  subject { described_class.new }

  describe '#category' do
    before { allow(Category).to receive(:new).and_return(:category) }

    its(:category) { should eq :category }
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    its(:call) { should be_success }

    its('call.object') { should eq category }
  end
end
