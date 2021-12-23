# frozen_string_literal: true

RSpec.describe Backoffice::Tags::InitializeService do
  subject { described_class.new category }

  let(:category) { build :category }

  describe '#tag' do
    before { allow(category).to receive_message_chain(:tags, :new).and_return(:tag) }

    its(:tag) { should eq :tag }
  end

  describe '#call' do
    let(:tag) { build :tag }

    before { allow(subject).to receive(:tag).and_return(tag) }

    its(:call) { should be_success }

    its('call.object') { should eq tag }
  end
end
