# frozen_string_literal: true

module Structure
  module List
    # SinglyNode class
    class SinglyNode
      def initialize(value = nil)
        @value = value
        @next = nil
      end
    end

    # Singly linked list class
    class Singly
      attr_reader :size

      def initialize
        @head = nil
        @tail = nil
        @size = 0
      end

      def empty?
        @size.zero?
      end

      def append(value = nil)
        new_node = SinglyNode.new(value)
        if empty?
          @head = new_node
        else
          @tail&.next = new_node
        end
        @tail = new_node
        @size += 1
      end

      def prepend(value = nil)
        new_node = SinglyNode.new(value)
        if empty?
          @tail = new_node
        else
          new_node.next = @head
        end
        @head = new_node
        @size += 1
      end

      def destroy
        @head = nil
        @tail = nil
        @size = 0
      end

      def delete_first
        return nil if empty?

        value = @head&.value
        @tail = nil if @size == 1
        @head = @head&.next
        @size -= 1
        value
      end

      def delete_last
        return nil if empty?

        return delete_first if @size == 1

        value = @tail&.value
        current = @head
        current = current&.next until current&.next&.next.nil?
        current&.next = nil
        @tail = current
        @size -= 1
        value
      end
    end
  end
end
