# frozen_string_literal: true

RSpec.describe Tags::GetCollectionService do
  subject { described_class.new params }

  let(:params) { { category_id: 15 } }

  describe '#category' do
    context do
      before { subject.instance_variable_set :@category, :category }

      its(:category) { should eq :category }
    end

    context do
      before { allow(Category).to receive(:find).with(15).and_return(:category) }

      its(:category) { should eq :category }
    end
  end

  describe '#call' do
    before { allow(subject).to receive_message_chain(:category, :tags, :order).with(:name).and_return(:tags) }

    its(:call) { should eq :tags }
  end
end
