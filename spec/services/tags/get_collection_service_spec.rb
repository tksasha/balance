# frozen_string_literal: true

RSpec.describe Tags::GetCollectionService do
  subject { described_class.new category }

  let(:category) { build :category }

  describe '#call' do
    before { allow(category).to receive_message_chain(:tags, :order).with(:name).and_return(:tags) }

    its(:call) { should eq :tags }
  end
end
